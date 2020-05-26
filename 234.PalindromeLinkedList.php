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
        $len = count($head);
        foreach($head as $key => $val){
            if($val != $head[$len - $key - 1]){
                return false;
            }
        }
        return true;
    }
}

$str = new Solution();
var_dump($str->isPalindrome([1]));
