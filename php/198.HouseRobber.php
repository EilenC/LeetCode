<?php
class Solution
{
    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function rob($nums)
    {
        $len = count($nums);
        if ($len == 0) return 0;
        $dp[0] = 0;
        $dp[1] = $nums[0];
        for ($i = 2; $i < $len + 1; $i++) {
            $dp[$i] = max($dp[$i - 1], $dp[$i - 2] + $nums[$i - 1]);
        }

        return $dp[$len];
    }
}
$str = new Solution();
var_dump($str->rob([0]));
