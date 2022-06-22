# FROM to specify the base image
# use alphine version to produce small output image
FROM golang:1.18.3-alpine3.16
# WORKDIR to declare the current working dir inside image
WORKDIR /app
# COPY files/folders from local to container
# first . indicates every files/folders in current dir
# everything under userAPI folder will be copied
# second . indicates the current working dir inside image
# where files and folders are copied to
# in this case will be /app
COPY . .
# RUN to build our app into single binary executable file
# -o means output, main is the name of output binary file
# then indicate main entrypoint of our app which is main.go
RUN go build -o main main.go
# best practice use EXPOSE to inform docker
# that the container listens on the specified network port
EXPOSE 5000
# CMD to define the command to run when container starts
# run our exe file which is the main
# where it is located inside app folder
CMD [ "/app/main" ]