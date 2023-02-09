firebase auth:export save_file.csv --format=csv --project $1
go run splitclean.go