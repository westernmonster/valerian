packr
swag init
gox -osarch="linux/amd64" -output="api"
rsync -avz --delete ./api root@118.31.11.232:/home/www/api/
rsync -avz --delete    ./docs/swagger/swagger.json root@118.31.11.232:/home/www/apidoc/build/
