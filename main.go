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
	log.Printf("writing ... (bucket: %s, name: %s)", bucket, name)
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("error creating client: %+v\n", err)
		return
	}

	acls, err := client.Bucket(bucket).ACL().List(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("new object acls: %v", acls)

	obj := client.Bucket(bucket).Object(name)
	wc := obj.NewWriter(ctx)
	wc.ContentType = "text/plain"

	publicRead := storage.ACLRule{storage.AllUsers, storage.RoleReader}

	wc.ACL = append(acls, publicRead)
	//wc.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}

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
	/*
		obj := client.Bucket(bucket).Object(filename)
		acls, err := obj.ACL().List(ctx)
		if err != nil {
			panic(err)
		}
	*/

}

func acls(bucket, filename string) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	obj := client.Bucket(bucket).Object(filename)
	acls, err := obj.ACL().List(ctx)
	if err != nil {
		panic(err)
	}

	for _, rule := range acls {
		fmt.Printf("rule (%T) %s has role %s\n", rule, rule.Entity, rule.Role)
	}

}

// read whats in the bucket
func read(bucket string) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: handle error.
	}
	defer client.Close()

	obj := client.Bucket(bucket).Object("test-file.txt")

	rc, err := obj.NewReader(ctx)
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

	// read("guler-bucket")
	//acls("guler-bucket", "test-file.txt")
	//	ls("guler-bucket")

	//write("guler-bucket", "lob.txt", "always look on the bright side of life")
	write("guler-bucket", "lob-public.txt", "always look on the bright side of life")

	acls("guler-bucket", "lob-public.txt")

}
