pipeline {
    agent any 
    stages {
        stage('Stage 1') {
            steps {
                @echo off
                IF not exist backup/%date:~10,4%%date:~4,2%%date:~7,2% (mkdir "backup/%date:~10,4%%date:~4,2%%date:~7,2%" && echo %date:~10,4%%date:~4,2%%date:~7,2%) ELSE (
                set a=0
                :_loop
                set /a a+=1
                IF not exist backup/%date:~10,4%%date:~4,2%%date:~7,2%^(%a%^) (mkdir "backup/%date:~10,4%%date:~4,2%%date:~7,2%(%a%)" && echo backup/%date:~10,4%%date:~4,2%%date:~7,2%^(%a%^) && goto _break)
                IF %a% GTR 10 (echo error && goto _break) ELSE (goto _loop)
                )
                :_break
                @echo on
            }
        }
    }
}
