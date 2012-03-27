// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

var enc64, dec64;
(function() {
	// modified base64url
	var ord = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_";
	var table = new Array(256);
	for(var i = 0; i < ord.length; i++)
		table[ord.charCodeAt(i)] = i;

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
			c = table[s.charCodeAt(i)];
			if(c === undefined)
				return -1;
			out = (out << 6) | c;
		}
		return out;
	};
})();
