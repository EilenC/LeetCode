<?php
class Solution
{

    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    function threeSum($nums)
    {
        $res = [];

        if (!$nums || ($len = count($nums)) < 3) {
            return $res;
        }

        sort($nums);
        foreach ($nums as $key => $val) {
            $left = $key + 1;
            $right = $len - 1;

            if ($val > 0) {
                break;
            }

            if ($key > 0 && $val === $nums[$key - 1]) {
                continue;
            }

            while ($left < $right) {
                $sum = $val + $nums[$left] + $nums[$right];

                if ($sum === 0) {
                    array_push($res, [$val, $nums[$left], $nums[$right]]);

                    while ($left < $right && $nums[$left] === $nums[++$left]);
                    while ($left < $right && $nums[$right] === $nums[--$right]);
                } else if ($sum < 0) {
                    $left++;
                } else {
                    $right--;
                }
            }
        }
        return $res;
    }
}
