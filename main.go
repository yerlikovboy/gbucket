package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"

	// Imports the Google Cloud Storage client package
	"cloud.google.com/go/storage"
)

func ls(name string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: handle error.
	}
	it := client.Bucket(name)
	fmt.Printf("bucket info: %+v\n", it)
	/*
		for {
			bucketAttrs, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				// TODO: Handle error.
			}
			fmt.Println(bucketAttrs)
		}
	*/
}

func write(bucket string, name string, content string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("error creating client: %+v\n", err)
		return
	}

	wc := client.Bucket(bucket).Object(name).NewWriter(ctx)

	// Write some text to obj. This will overwrite whatever is there.
	if _, err := fmt.Fprintf(wc, content); err != nil {
		log.Printf("error writing to bucket: %+v\n", err)
		return
	}
	// Close, just like writing a file.
	if err := wc.Close(); err != nil {
		log.Printf("error closing handle: %+v\n")
		return
	}
}

// read whats in the bucket
func read(bucket string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: handle error.
	}
	rc, err := client.Bucket(bucket).Object("test-file.txt").NewReader(ctx)
	if err != nil {
		log.Printf("error opening reader: %+v\n", err)
		return
	}
	slurp, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		log.Printf("error reading file contents: %+v\n", err)
		return
	}
	fmt.Println("file contents:", string(slurp))
}

func create(name string) error {
	ctx := context.Background()

	// Your Google Cloud Platform project ID
	projectID := "594087932715"

	// Creates a client
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// The name for the new bucket
	bucketName := "alis-new-bucket"

	// Prepares a new bucket
	bucket := client.Bucket(bucketName)

	// Creates the new bucket
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	fmt.Printf("Bucket %v created.", bucketName)
	return nil
}

// export JSON with application key to GOOGLE_APPLICATION_CREDENTIALS
func main() {
	/*
		ctx := context.Background()

		// Your Google Cloud Platform project ID
		projectID := "594087932715"

		// Creates a client
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		// The name for the new bucket
		bucketName := "alis-new-bucket"

		// Prepares a new bucket
		bucket := client.Bucket(bucketName)

		// Creates the new bucket
		if err := bucket.Create(ctx, projectID, nil); err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
		}

		fmt.Printf("Bucket %v created.", bucketName)
	*/

	read("guler-bucket")
	ls("guler-bucket")
	write("guler-bucket", "lob.txt", "always look on the bright side of life")

}
