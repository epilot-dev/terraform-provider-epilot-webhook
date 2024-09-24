// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type S3Ref struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

func (o *S3Ref) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *S3Ref) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}
