<?php
class Solution
{

    /**
     * @param String $s
     * @return Boolean
     */
    function isValid($s)
    {
        if (strlen($s) === 0) return true;
        if (strlen($s) % 2 != 0) return false;
        $s = str_split($s, 1);
        $stack = [];
        foreach ($s as $key => $val) {
            if ($val === ')' || $val === '}' || $val === ']') {
                if (!empty($stack) && end($stack) === $val) {
                    array_pop($stack);
                } else {
                    return false;
                }
            } else {
                if ($val === '(') {
                    $stack[] = ')';
                } else if ($val === '[') {
                    $stack[] = ']';
                } else if ($val === '{') {
                    $stack[] = '}';
                }
            }
        }
        return empty($stack);
    }
}

$str = new Solution();
var_dump($str->isValid("(]"));
