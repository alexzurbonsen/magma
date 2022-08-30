/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import axios from 'axios';
import bcrypt from 'bcryptjs';
import {AccessRoles} from '../shared/roles';
import {User} from '../shared/sequelize_models';
import {UserModel} from '../shared/sequelize_models/models/user';
import {
  createJoinedOrganization,
  getOrganizationByName,
} from '../server/host/routes';

const SALT_GEN_ROUNDS = 10;

type UserObject = {
  organization: string;
  email: string;
  password: string;
  role: number;
};

async function updateUser(user: UserModel, userObject: UserObject) {
  const {password, role} = userObject;
  const salt = await bcrypt.genSalt(SALT_GEN_ROUNDS);
  const passwordHash = await bcrypt.hash(password, salt);
  await user.update({
    password: passwordHash,
    role,
  });
}

async function createUser(userObject: UserObject) {
  const {organization, email, password, role} = userObject;
  const salt = await bcrypt.genSalt(SALT_GEN_ROUNDS);
  const passwordHash = await bcrypt.hash(password, salt);
  const org = await createOrFetchOrganization(organization);
  await User.create({
    email: email.toLowerCase(),
    password: passwordHash,
    role,
    networkIDs: [],
    organization: org,
  });
}

async function createOrUpdateUser(userObject: UserObject) {
  const user = await User.findOne({
    where: {
      organization: userObject.organization,
      email: userObject.email.toLowerCase(),
    },
  });
  if (!user) {
    await createUser(userObject);
  } else {
    await updateUser(user, userObject);
  }
}

async function createOrFetchOrganization(
  organization: string,
): Promise<string> {
  const joinedOrganization = await getOrganizationByName(organization);
  if (!joinedOrganization) {
    console.log(`Creating a new Organization: name=${organization}`);
    const createdOrganization = await createJoinedOrganization(
      organization,
      [],
      [],
    );
    return createdOrganization.name;
  } else {
    return joinedOrganization.name;
  }
}

function main() {
  const args = process.argv.slice(2);
  if (args.length !== 3) {
    console.log('Usage: setPassword.ts <organization> <email> <password>');
    process.exit(1);
  }
  const userObject = {
    organization: args[0],
    email: args[1],
    password: args[2],
    role: AccessRoles.SUPERUSER,
  };
  console.log('Creating a new user: email=' + userObject.email);
  createOrUpdateUser(userObject)
    .then(() => {
      console.log('Success');
      process.exit();
    })
    .catch(err => {
      if (axios.isAxiosError(err)) {
        console.log(
          `Error: Status: ${
            err?.response?.status ?? 500
          }: ${(err as Error).toString()}`,
        );
      } else {
        console.error(err);
      }
      process.exit(1);
    });
}

main();
