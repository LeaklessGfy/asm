package asm

// FieldVisitor a visitor to visit a Java field. The methods of this class must be called in the following order:
// ( <tt>visitAnnotation</tt> | <tt>visitTypeAnnotation</tt> | <tt>visitAttribute</tt> )*
// <tt>visitEnd</tt>.
type FieldVisitor interface {
	VisitAnnotation(descriptor string, visible bool) AnnotationVisitor
	VisitTypeAnnotation(typeRef int, typePath *TypePath, descriptor string, visible bool) AnnotationVisitor
	VisitAttribute(attribute *Attribute)
	VisitEnd()
}
