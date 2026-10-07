<?php
require_once('plugins/login-password-less.php');

// SQLite has no users, so Adminer itself verifies this password.
// Hash of "admin", regenerate with:
//   docker run --rm --entrypoint php adminer:latest -r "echo password_hash('YOUR_PASSWORD', PASSWORD_DEFAULT), PHP_EOL;"
return new AdminerLoginPasswordLess('$2y$12$pxMjmRfIhgTDiw0T5IvPaOFp3HR5TNKSG4djVP.ymNQiyTbSSX9t.');
