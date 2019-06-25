Feature: command-line parser
  In order to have a reliable parser
  As a test suite
  I need to snapshot test the arguments

  Scenario: should correctly parse the arguments
    Given the arguments:
    """
    git clone -v -q --progress -n --bare --mirror -l --no-hardlinks -s --recurse-submodules=path/spec -j=10 --template=template/directory --reference="git@github.com:celicoo/docli.git" --dissociate -o=name -b=branch -u=path --depth=1 --shallow-since=2019-16-05 --single-branch --no-tags --shallow-submodules -4 -6
    """ 
    When I call the parser
    Then the returning value should be equal
    """
    {[{git { }} {clone { }} {-v { }} {-q { }} {--progress { }} {-n { }} {--bare { }} {--mirror { }} {-l { }} {--no-hardlinks { }} {-s { }} {--recurse-submodules {= path/spec}} {-j {= 10}} {--template {= template/directory}} {--reference {= "git@github.com:celicoo/docli.git"}} {--dissociate { }} {-o {= name}} {-b {= branch}} {-u {= path}} {--depth {= 1}} {--shallow-since {= 2019-16-05}} {--single-branch { }} {--no-tags { }} {--shallow-submodules { }} {-4 { }} {-6 { }}] }
    """