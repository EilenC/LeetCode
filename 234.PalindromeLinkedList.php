<?php
// Definition for a singly-linked list.
// class ListNode
// {
//     public $val = 0;
//     public $next = null;
//     function __construct($val)
//     {
//         $this->val = $val;
//     }
// }
class Solution
{

    /**
     * @param ListNode $head
     * @return Boolean
     */
    function isPalindrome($head)
    {
        while ($head->next != null || $head->val != null) {
            $array[] = $head->val;
            $head = $head->next;
        }
        unset($head);
        $len = count($array);
        foreach ($array as $key => $val) {
            if ($val != $array[$len - $key - 1]) {
                return false;
            }
        }
        return true;
    }
}

$str = new Solution();
var_dump($str->isPalindrome([1, 2]));
