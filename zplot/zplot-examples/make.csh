#!/bin/csh -f
# Python 2.7.18 :: Anaconda, Inc.

@ count = 1
foreach dir ( * )
    if (-d $dir) then
        # echo $dir
        # echo $*
        cd $dir
        rm -rf *.eps >& /dev/null
        rm -rf *.svg >& /dev/null
        rm -rf *.pdf >& /dev/null
        foreach type ( pdf )
            foreach py ( *.py )
                echo "$count ${dir}/${py} $type"
                @ count += 1
                python ${py} $type
            end
        end
        mv *.pdf ..
        mv *.eps ..
        mv *.svg ..
        cd ..
    endif
end
