# Use the official Node.js image as the base image
FROM node:14-alpine as build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the package.json and package-lock.json files and install dependencies
COPY package*.json ./
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the React app
RUN npm run build

# Use the official nginx image to serve the React app
FROM nginx:alpine

# Copy the built React app to the nginx html directory
COPY --from=build /app/build /usr/share/nginx/html

# Expose port 80 to the outside world
EXPOSE 80

# Command to run nginx
CMD ["nginx", "-g", "daemon off;"]
