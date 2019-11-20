docker build -t eferhatg/cas-server ./server
docker push eferhatg/cas-server 
docker build -t eferhatg/cas-cache ./cacheserver
docker push eferhatg/cas-cache 
docker build -t eferhatg/cas-client ./client
docker push eferhatg/cas-client 

# docker build -t eferhatg/cas-server ./server
# docker build -t eferhatg/cas-client ./client


# docker push eferhatg/cas-client 
# docker push eferhatg/cas-server 
# docker push eferhatg/cas-cache 