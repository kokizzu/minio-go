package minio

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/kokizzu/minio-go/v7/pkg/s3utils"
)

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketCors.html

const PutBucketCorsXmlExample = `<?xml version="1.0" encoding="UTF-8"?>
<CORSConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
  <CORSRule>
      <AllowedOrigin>*</AllowedOrigin>

      <AllowedMethod>GET</AllowedMethod>
      <AllowedMethod>PUT</AllowedMethod>
      <AllowedMethod>POST</AllowedMethod>
      <AllowedMethod>DELETE</AllowedMethod>

      <AllowedHeader>*</AllowedHeader>

      <MaxAgeSeconds>3000</MaxAgeSeconds>

      <ExposeHeader>x-amz-server-side-encryption</ExposeHeader>
  </CORSRule>
</CORSConfiguration>`

// PutBucketCors set bucket CORS policy
func (c *Client) PutBucketCors(ctx context.Context, bucketName string, xmlCors string) (err error) {
	// Validate the input arguments.
	if err := s3utils.CheckValidBucketNameStrict(bucketName); err != nil {
		return err
	}

	urlValues := make(url.Values)
	urlValues.Set(`cors`, ``)

	// PUT bucket request metadata.
	reqMetadata := requestMetadata{
		bucketName:  bucketName,
		contentBody: strings.NewReader(xmlCors),
		queryValues: urlValues,
	}

	// Execute PUT to create a new bucket.
	resp, err := c.executeMethod(ctx, http.MethodPut, reqMetadata)
	defer closeResponse(resp)
	if err != nil {
		return err
	}

	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			return httpRespToErrorResponse(resp, bucketName, "")
		}
	}

	// Success.
	return nil
}

// GetBucketCors get bucket CORS policy
func (c *Client) GetBucketCors(ctx context.Context, bucketName string) (xmlCors string, err error) {
	// Validate the input arguments.
	if err := s3utils.CheckValidBucketNameStrict(bucketName); err != nil {
		return ``, err
	}

	urlValues := make(url.Values)
	urlValues.Set(`cors`, ``)

	// PUT bucket request metadata.
	reqMetadata := requestMetadata{
		bucketName:  bucketName,
		queryValues: urlValues,
	}

	// Execute PUT to create a new bucket.
	resp, err := c.executeMethod(ctx, http.MethodGet, reqMetadata)
	defer closeResponse(resp)
	if err != nil {
		return ``, err
	}

	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			return ``, httpRespToErrorResponse(resp, bucketName, "")
		}
	}

	var xmlCorsByt []byte
	xmlCorsByt, err = io.ReadAll(resp.Body)

	// Success.
	return string(xmlCorsByt), nil
}
