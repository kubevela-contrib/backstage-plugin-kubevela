package main

// Entity is the type of backstage entity
type Entity struct {
	ApiVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Metadata   *EntityMeta            `json:"metadata"`
	Spec       map[string]interface{} `json:"spec"`
	Relations  []EntityRelation       `json:"relations,omitempty"`
}

type VelaBackstageTrait struct {
	Type        string            `json:"type,omitempty"`
	LifeCycle   string            `json:"lifecycle,omitempty"`
	Owner       string            `json:"owner,omitempty"`
	System      string            `json:"system,omitempty"`
	Description string            `json:"description,omitempty"`
	Title       string            `json:"title,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Links       []EntityLink      `json:"links,omitempty"`
	Targets     []string          `json:"targets,omitempty"`
}

type EntityMeta struct {
	/**
	 * A globally unique ID for the entity.
	 *
	 * This field can not be set by the user at creation time, and the server
	 * will reject an attempt to do so. The field will be populated in read
	 * operations. The field can (optionally) be specified when performing
	 * update or delete operations, but the server is free to reject requests
	 * that do so in such a way that it breaks semantics.
	 */
	Uid string `json:"uid,omitempty"`
	/**
	 * An opaque string that changes for each update operation to any part of
	 * the entity, including metadata.
	 *
	 * This field can not be set by the user at creation time, and the server
	 * will reject an attempt to do so. The field will be populated in read
	 * operations. The field can (optionally) be specified when performing
	 * update or delete operations, and the server will then reject the
	 * operation if it does not match the current stored value.
	 */
	Etag string `json:"etag,omitempty"`

	Name string `json:"name"`
	/**
	 * The namespace that the entity belongs to.
	 */
	Namespace string `json:"namespace,omitempty"`
	/**
	 * A display name of the entity, to be presented in user interfaces instead
	 * of the `name` property above, when available.
	 *
	 * This field is sometimes useful when the `name` is cumbersome or ends up
	 * being perceived as overly technical. The title generally does not have
	 * as stringent format requirements on it, so it may contain special
	 * characters and be more explanatory. Do keep it very short though, and
	 * avoid situations where a title can be confused with the name of another
	 * entity, or where two entities share a title.
	 *
	 * Note that this is only for display purposes, and may be ignored by some
	 * parts of the code. Entity references still always make use of the `name`
	 * property, not the title.
	 */
	Title string `json:"title,omitempty"`
	/**
	 * A short (typically relatively few words, on one line) description of the
	 * entity.
	 */
	Description string `json:"description,omitempty"`
	/**
	 * Key/value pairs of identifying information attached to the entity.
	 */
	Labels map[string]string `json:"labels,omitempty"`
	/**
	 * Key/value pairs of non-identifying auxiliary information attached to the
	 * entity.
	 */
	Annotations map[string]string `json:"annotations,omitempty"`
	/**
	 * A list of single-valued strings, to for example classify catalog entities in
	 * various ways.
	 */
	Tags []string `json:"tags,omitempty"`
	/**
	 * A list of external hyperlinks related to the entity.
	 */
	Links []EntityLink `json:"links,omitempty"`
}

type EntityLink struct {

	/**
	 * The url to the external site, document, etc.
	 */
	Url string `json:"url"`
	/**
	 * An optional descriptive title for the link.
	 */
	Title string `json:"title,omitempty"`
	/**
	 * An optional semantic key that represents a visual icon.
	 */
	Icon string `json:"icon,omitempty"`
	/**
	 * An optional value to categorize links into specific groups
	 */
	Type string `json:"type,omitempty"`
}

type EntityRelation struct {
	/**
	 * The type of the relation.
	 */
	Type string `json:"type"`
	/**
	 * The entity ref of the target of this relation.
	 */
	TargetRef string `json:"targetRef"`
}
