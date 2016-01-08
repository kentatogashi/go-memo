use strict;
use warnings;
use Parallel::ForkManager;


my $pm = Parallel::ForkManager->new(500);
my @list1 = ();
for(my $i = 0; $i < 10000; $i++) {
  push @list1, $i;
}
foreach my $val (@list1) {
$pm->start and next;
srand;
sleep rand(3);
print $$ . ":" . $val . "\n";
$pm->finish;
}
$pm->wait_all_children();
