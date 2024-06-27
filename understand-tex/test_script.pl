use strict;
use warnings;
use Test::More tests => 5;

my $x = 2 + 2;

is($x, 4, '2 + 2 should equal 4');
ok($x == 4, 'x should be 4');
isnt($x, 5, 'x should not be 5');
like("Hello, world!", qr/world/, 'string should contain "world"');
unlike("Hello, world!", qr/planet/, 'string should not contain "planet"');
