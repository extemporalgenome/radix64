// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

var enc64, dec64;
(function() {
	// base64url
	var ord = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_";
	var dmin = '0'.charCodeAt(0), dmax = '9'.charCodeAt(0);
	var umin = 'A'.charCodeAt(0), umax = 'Z'.charCodeAt(0), uoff = ord.indexOf('A');
	var lmin = 'a'.charCodeAt(0), lmax = 'z'.charCodeAt(0), loff = ord.indexOf('a');
	var c62 = ord.charCodeAt(62), c63 = ord.charCodeAt(63);
	enc64 = function(n, len) {
		var out = []; 
		while(n > 0 || len > 0) {
			out.unshift(ord[n & 63]);
			n >>= 6;
			len--;
		}
		return out.join("");
	};
	dec64 = function(s) {
		var c, out = 0; 
		for(var i = 0; i < s.length; i++) {
			c = s.charCodeAt(i);
			if(c >= dmin && c <= dmax)
				c -= dmin;
			else if(c >= umin && c <= umax)
				c = c - umin + uoff;
			else if(c >= lmin && c <= lmax)
				c = c - lmin + loff;
			else if(c == c62)
				c = 62;
			else if(c == c63)
				c = 63;
			else
				return -1;
			out = (out << 6) | c;
		}
		return out;
	};
})();
