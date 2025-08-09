/* Bootstrap projects with created template */
package mraziz

/* Project item that can be generated */
type Generatable interface {
	Generate() error
}
