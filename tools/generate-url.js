// Import Google Cloud Storage and path modules
const { Storage } = require('@google-cloud/storage');
const path = require('path');

// Create a Storage client using the service account key
const storage = new Storage({
  keyFilename: path.join(__dirname, 'project-storege-key.json'),
});

// Define the name of the GCS bucket
const bucketName = 'project-storege';
// Generate a unique file name with timestamp
const fileName = `upload-${Date.now()}.jpg`; 

// Function to generate a signed URL for uploading a file
async function generateSignedUrl() {
  const options = {
    version: 'v4',
    action: 'write',
    expires: Date.now() + 7 * 24 * 60 * 60 * 1000, // 7 days code
    contentType: 'application/octet-stream',
  };

    // Generate the signed URL using GCS client
    const [url] = await storage
      .bucket(bucketName)
      .file(fileName)
      .getSignedUrl(options);

    // Output the signed URL for use in the frontend uploader
    console.log('\n Wklej ten URL do components/FileUploader:\n');
    console.log(url);
}
// Run the function
generateSignedUrl();
