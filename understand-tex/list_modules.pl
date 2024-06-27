#!/usr/bin/env perl

use strict;
use warnings;
use ExtUtils::Installed;

my $inst = ExtUtils::Installed->new();
my @modules = $inst->modules();

foreach my $module (@modules) {
    my $version = $inst->version($module) || 'unknown';
    print "$module - $version\n";
}

