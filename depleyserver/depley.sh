kill -9 $(pgrep webserve)
cd ~/shortbook-app/
git pull https://github.com/huangrongfa/Shortbook-app.git
cd webserve
./webserve &