<?php

$data=array(
    'category'=>'this is category',
    'article'=>array(
        array(
            'title'=>'title one xxxxxxxxxxxxx',
            'author'=>'xxxxx'
            ),
        array(
            'title'=>'title two yyyyyyyyyyyyyyy',
            'author'=>'yyyyy'
            )
        )
    );
echo json_encode($data);

