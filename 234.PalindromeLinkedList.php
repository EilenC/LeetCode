<?php
// Definition for a singly-linked list.
class ListNode
{
    public $val = 0;
    public $next = null;
    function __construct($val)
    {
        $this->val = $val;
    }
}
class Solution
{

    /**
     * @param ListNode $head
     * @return Boolean
     */
    // 翻转数组对比
    // function isPalindrome($head)
    // {
    //     while ($head->next != null || $head->val != null) {
    //         $array[] = $head->val;
    //         $head = $head->next;
    //     }
    //     unset($head);
    //     $len = count($array);
    //     foreach ($array as $key => $val) {
    //         if ($val != $array[$len - $key - 1]) {
    //             return false;
    //         }
    //     }
    //     return true;
    // }

    // 堆栈解法
    function isPalindrome($head)
    {
        $p = $head;
        if($head->next == null){
            return true;
        }
        while($head != null){
            $stack[]  = $head->val;
            $head = $head->next;
        }

        while($p != null){
            if($p->val != $stack[count($stack)-1]) return false;
            array_pop($stack);
            $p = $p->next;
        }
        return true;
    }
}

$str = new Solution();
var_dump($str->isPalindrome([-129,-129]));
