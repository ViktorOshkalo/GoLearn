# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build your app
RUN go build -o main .

# Expose the port your app will run on
EXPOSE 50051

# Command to run your application
CMD ["./main"]