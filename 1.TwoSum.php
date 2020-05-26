<?php
class Solution
{

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target)
    {
        $map = [];
        foreach ($nums as $key => $val) {

            $dif = $target - $val;

            if (!isset($map[$dif])) {
                $map[$val] = $key;
                continue;
            }
            return [$map[$dif], $key];
        }
    }
}
