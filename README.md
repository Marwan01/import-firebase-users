# import-firebase-users

### Dependencies

- Go
- firebase-cli

### Usage

Login to firebase-cli:
```
firebase login
```
Get your project ID:
```
firebase projects:list
```
Make script executable:
```
chmod +x entrypoint.sh
```
Finally run the script:
```
./entrypoint.sh <firebase project ID>
```