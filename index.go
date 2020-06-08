package gaia

// IndexToggleOpen can toggle the
// open/close status of an index
// Args:
// - s string
// 	the state of the index (open, close)
// - i string
//  the name of the index to open/close
func (c *Client) IndexToggleOpen(s string, i string) string {
	r := c.newRequest()
	return r.post(i+"/_"+s, "")
}

// IndexDelete deletes an index
// Use with caution
// Args:
// - i string
//  the name of the index to delete
func (c *Client) IndexDelete(i string) string {
	r := c.newRequest()
	return r.delete(i)
}

// IndexGet retrieves information about one or more indexes
// Args:
// - i string
//  comma seperated list of indices to retrieve, or _all
func (c *Client) IndexGet(i string) string {
	r := c.newRequest()
	return r.get(i)
}

// IndexAliasesGet filters information about an index to only
// include the specific alias features
// Args:
// - i string
//  comma seperated list of indices to retrieve, or _all
func (c *Client) IndexAliasesGet(i string) string {
	r := c.newRequest()
	return r.get(i + "/_aliases")
}

// IndexSettingsGet filters information about an index to only
// include the specific settings features
// Args:
// - i string
//  comma seperated list of indices to retrieve, or _all
func (c *Client) IndexSettingsGet(i string) string {
	r := c.newRequest()
	return r.get(i + "/_settings")
}

// IndexMappingsGet filters information about an index to only
// include the specific alias features
// Args:
// - i string
//  comma seperated list of indices to retrieve, or _all
func (c *Client) IndexMappingsGet(i string) string {
	r := c.newRequest()
	return r.get(i + "/_mappings")
}
