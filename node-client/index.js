const path = require('path');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const proto = protoLoader.loadSync(path.join(__dirname, 'posts_service.proto'));
const definition = grpc.loadPackageDefinition(proto);

const serverUrl = 'localhost:10000';

const client = new definition.PostService(serverUrl, grpc.credentials.createInsecure())
    
client.getPosts({}, (error, posts) => {
    if (!error) {
        console.log(posts)
    } else {
        console.error(error)
    }
})