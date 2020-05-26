<?php
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function lengthOfLongestSubstring($s) {
        $n = strlen($s);
        if ($n <= 1) return $n;
        $max = 0;
        $hash = [];
        $left = 0;
        for ($i = 0; $i < $n; ++$i) {
            $char = substr($s, $i, 1);
            if (isset($hash[$char])) {
                $left = max($left, $hash[$char] + 1);
            }
            $hash[$char] = $i;
            $max = max($max, $i - $left + 1);
        }
        return $max;
    }
}
$str = new Solution();
echo $str->lengthOfLongestSubstring("abcabcbbefg"); 