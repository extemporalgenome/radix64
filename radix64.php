<?php

$_ord = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_";
$_table = array_flip(str_split($_ord));

function encode($n, $len = 0) {
    global $_ord;
    $out = "";
    while ($n > 0 || $len > 0) {
        $out = $_ord[$n & 63] . $out;
        $n >>= 6;
        $len--;
    }
    $out = str_pad($out, 10, "0", STR_PAD_LEFT);
    return $out;
}

function decode($input) {
    global $_table;
    $n = 0;
    foreach (str_split($input) as $c) {
        $c = $_table[$c];
        if ($c === null) {
            throw new Exception("Invalid character in input: " . $c);
        }
        $n = ($n << 6) | $c;
    }
    return $n;
}

// Example usage
$encoded = encode("123456789");
echo "Encoded: " . $encoded . "\n";

$decoded = decode($encoded);
echo "Decoded: " . $decoded . "\n";

?>
