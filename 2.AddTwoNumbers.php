<?php
// Definition for a singly-linked list.
 class ListNode {
     public $val = 0;
     public $next = null;
     function __construct($val) { $this->val = $val; }
 }
class Solution
{
    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2)
    {
        $sum = 0;
        $nextSum = 0;
        $res = new ListNode(0);
        $resTemp = $res;

        while ($l1 != null && $l2 != null) {
            $sum = $l1->val + $l2->val + $nextSum;
            if ($sum >=  10) {
                $resTemp->next = new ListNode($sum - 10);
                $nextSum = 1;
            } else {
                $resTemp->next = new ListNode($sum);
                $nextSum = 0;
            }
            $resTemp = $resTemp->next;
            $l1 = $l1->next;
            $l2 = $l2->next;
        }

        while ($l1 != null) {
            $sum = $l1->val + $nextSum;
            if ($sum >= 10) {
                $resTemp->next = new ListNode($sum - 10);
                $nextSum = 1;
            } else {
                $l1->val = $sum;
                $resTemp->next = $l1;
                $nextSum = 0;
                break;
            }
            $resTemp = $resTemp->next;
            $l1 = $l1->next;
        }

        while ($l2 != null) {
            $sum = $l2->val + $nextSum;
            if ($sum >= 10) {
                $resTemp->next = new ListNode($sum - 10);
                $nextSum = 1;
            } else {
                $l2->val = $sum;
                $resTemp->next = $l2;
                $nextSum = 0;
                break;
            }
            $resTemp = $resTemp->next;
            $l2 = $l2->next;
        }

        if ($nextSum == 1) {
            $resTemp->next = new ListNode(1);
        }
        return $res->next;
    }
}
