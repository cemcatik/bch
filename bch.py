import bcrypt
import argparse
import getpass

def out(x):
    print(x.decode('utf-8'))

parser = argparse.ArgumentParser(description='Generate and verify bcrypt password hashes')
parser.add_argument('password', nargs='?', help='password to hash')
parser.add_argument('-f', '--factor', type=int, required=False, default=12, help='work factor', metavar='F')
parser.add_argument('-v', '--verify', type=str, required=False, help='hash to verify against entered password', metavar='hash')

args = parser.parse_args()

def password():
    if args.password:
        return args.password.encode('utf-8')
    else:
        return getpass.getpass()

if args.verify:
    hashed = args.verify
    match = bcrypt.hashpw(password(), hashed)
    print(match == hashed)
else:
    factor = args.factor
    salt = bcrypt.gensalt(factor)
    hash = bcrypt.hashpw(password(), salt)
    out(hash)
