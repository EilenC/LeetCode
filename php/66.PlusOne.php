<?php
class Solution
{
    /**
     * @param Integer[] $digits
     * @return Integer[]
     */
    function plusOne($digits)
    {
        foreach ($digits = array_reverse($digits) as $key => $val) {
            if ($val == 9) {
                $digits[$key] = 0;
            } else {
                $digits[$key] += 1;
                break;
            }
        }
        $digits = array_reverse($digits);
        if ($digits[0] == 0) {
            $digits[0] = 1;
            $digits[] = 0;
        }
        return $digits;
    }
}
