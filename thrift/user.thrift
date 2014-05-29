include "shared.thrift"

/* remember to update generated package from 'generated/user' to 'source.discoveryeducation.com/users/mhatch/repos/de2/thrift/gen-go/generated/user'
* thrift automatically converts '.' to new folder subdirectories instead of creating a single folder with a '.' in the name */

namespace go generated.user


service UserServicer extends shared.BinaryMessenger {
}