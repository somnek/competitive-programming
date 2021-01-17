<?php
    function flip(string $dir, array $arr): array{
        $dir == 'R'?sort($arr): rsort($arr);
        return $arr;
    }
    print_r(flip('R', [6, 1, 2, 3, 4]));
?>