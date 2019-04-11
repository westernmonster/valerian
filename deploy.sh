packr
gox -osarch="linux/amd64" -output="api"
rsync -avz --delete ./api root@118.31.11.232:/home/www/api/
rsync -avz --delete ./migrations root@118.31.11.232:/home/www/api/
rsync -avz --delete ./docs root@118.31.11.232:/home/www/api/
rsync -avz --delete ./config.json root@118.31.11.232:/home/www/api/
rsync -avz --delete ./dbconfig.yml root@118.31.11.232:/home/www/api/
