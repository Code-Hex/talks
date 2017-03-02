package Animal {
    sub new {
        my $class = shift;
        bless { @_ }, $class;
    }
    sub say { print shift->{say}."\n" }
};
package Dog {
    @ISA = qw/Animal/;
    sub new { bless Animal->SUPER::new(say => "Waon!!"), shift }
};
package Cat {
    @ISA = qw/Animal/;
    sub new { bless Animal->SUPER::new(say => "Nyan!!"), shift }
};
my ($dog, $cat) = (Dog->new, Cat->new);
$dog->say();
$cat->say();
print $dog->isa('Animal'), ", ", $cat->isa('Animal');