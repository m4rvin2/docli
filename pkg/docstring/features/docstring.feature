Feature: docstring parser
  In order to have a reliable parser
  As a test suite
  I need to snapshot test the docstring

  Scenario: should correctly parse the docstring
    Given the docstring:
    """
    usage: git clone [<options>] [--] <repo> [<dir>]

      -v, --verbose                      be more verbose
      -q, --quiet                        be more quiet
      --progress                         force progress reporting
      -n, --no-checkout                  don't create a checkout
      --bare                             create a bare repository
      --mirror                           create a mirror repository (implies bare)
      -l, --local                        to clone from a local repository
      --no-hardlinks                     don't use local hardlinks, always copy
      -s, --shared                       setup as shared repository
      --recurse-submodules[=<pathspec>]  initialize submodules in the clone
      -j, --jobs <n>                     number of submodules cloned in parallel
      --template <template-directory>    directory from which templates will be used
      --reference <repo>                 reference repository
      --reference-if-able <repo>         reference repository
      --dissociate                       use --reference only while cloning
      -o, --origin <name>                use <name> instead of 'origin' to track upstream
      -b, --branch <branch>              checkout <branch> instead of the remote's HEAD
      -u, --upload-pack <path>           path to git-upload-pack on the remote
      --depth <depth>                    create a shallow clone of that depth
      --shallow-since <time>             create a shallow clone since a specific time
      --shallow-exclude <revision>       deepen history of shallow clone, excluding rev
      --single-branch                    clone only one branch, HEAD or --branch
      --no-tags                          don't clone any tags, and make later fetches not to follow them
      --shallow-submodules               any cloned submodules will be shallow
      --separate-git-dir <gitdir>        separate git dir from working tree
      -c, --config <key=value>           set config inside the new repository
      -4, --ipv4                         use IPv4 addresses only
      -6, --ipv6                         use IPv6 addresses only
      --filter <args>                    object filtering
    """ 
    When I call the parser
    Then the returning value should be equal
    """
    {[{
       [{-v , } {--verbose }]} {
       [{-q , } {--quiet }]} {
       [{--progress }]} {
       [{-n , } {--no-checkout }]} {
       [{--bare }]} {
       [{--mirror }]} {
       [{-l , } {--local }]} {
       [{--no-hardlinks }]} {
       [{-s , } {--shared }]} {
       [{--recurse-submodules }]} {
       [{-j , } {--jobs }]} {
       [{--template }]} {
       [{--reference }]} {
       [{--reference-if-able }]} {
       [{--dissociate }]} {
       [{-o , } {--origin }]} {
       [{-b , } {--branch }]} {
       [{-u , } {--upload-pack }]} {
       [{--depth }]} {
       [{--shallow-since }]} {
       [{--shallow-exclude }]} {
       [{--single-branch }]} {
       [{--no-tags }]} {
       [{--shallow-submodules }]} {
       [{--separate-git-dir }]} {
       [{-c , } {--config }]} {
       [{-4 , } {--ipv4 }]} {
       [{-6 , } {--ipv6 }]} {
       [{--filter }]}] }
    """