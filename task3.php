<?php

function countTuesdays($startDate, $endDate) {
    $start = new DateTime($startDate);
    $end = new DateTime($endDate);
    $end = $end->modify('+1 day'); // включаем конечную дату в расчет

    $totalDays = $start->diff($end)->days;

    $totalTuesdays = (int)($totalDays / 7);

    // Дополнительные вторники в зависимости от дня недели начальной и конечной даты
    if ($start->format('N') <= 2 && $end->format('N') >= 2) {
        $totalTuesdays++;
    }

    return $totalTuesdays;
}

$startDate = "2000-01-01";
$endDate = "2023-12-31";
echo countTuesdays($startDate, $endDate);