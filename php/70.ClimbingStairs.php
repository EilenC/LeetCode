<?php
class Solution
{

    /**
     * @param Integer $n
     * @return Integer
     */
    // 斐波那契计算
    // function climbStairs($n)
    // {
    //     $dp[1] = 1;
    //     $dp[2] = 2;
    //     for ($i = 3; $i <= $n; $i++) {
    //         $dp[$i] = $dp[$i - 1] + $dp[$i - 2];
    //     }

    //     return $dp[$n];
    // }
    // 斐波那契计算2
    function climbStairs($n)
    {
        static $cache = [1 => 1, 2 => 2];
        if ($n === 1) return 1;
        if ($n === 2) return 2;

        if (!isset($cache[$n - 1])) {
            $cache[$n - 1] = $this->climbStairs($n - 1);
        }
        if (!isset($cache[$n - 2])) {
            $cache[$n - 2] = $this->climbStairs($n - 2);
        }
        $cache[$n] = $cache[$n - 1] + $cache[$n - 2];
        return $cache[$n];
    }
}
$str = new Solution();
$str->climbStairs(5);
