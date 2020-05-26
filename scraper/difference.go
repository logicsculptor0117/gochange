/*
This file is under GNU AFFERO GENERAL PUBLIC LICENSE

Permissions of this strongest copyleft license are conditioned
on making available complete source code of licensed works and
modifications, which include larger works using a licensed work,
under the same license. Copyright and license notices must be preserved.
Contributors provide an express grant of patent rights.
When a modified version is used to provide a service over a network,
the complete source code of the modified version must be made available.

Edoardo Ottavianelli, https://edoardoottavianelli.it
 */

package scraper

import (
	"time"
)

// TODO
type Change struct {
	monitor 	Monitor
	body		string
	timestamp	int64
}

// TODO
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// TODO
func Edited(change Change, content string) bool {
	return change.body != content
}