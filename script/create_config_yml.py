#!/usr/bin/env python3

import sys

if __name__ == '__main__':
    if len(sys.argv) != 5:
            exit("use: python3 create_config_yml.py <number of households> <number of validators> <number of token> <value of stake>")

    households = sys.argv[1]
    token = sys.argv[3]
    stake = sys.argv[4]
    validators = sys.argv[2]

    file = open("../config.yml", "w")

    #build accounts
    file.write("accounts:"+"\n")
    for household in range(int(households)):
        file.write("  - name: " +"household"+str(household)+"\n")
        file.write('    coins: ["'+token+'token", "'+stake+'stake"]'+"\n")

    #build validators
    file.write("validator:"+"\n")
    for validator in range(int(validators)):
        file.write("  name: "+"household"+str(validator)+"\n")
        file.write('  staked: "'+stake+'stake"'+"\n")

    #build client
    file.write("client:"+"\n")
    file.write("  vuex:"+"\n")
    file.write('    path: "vue/src/store"'+"\n")
    file.write("  openapi:"+"\n")
    file.write('    path: "docs/static/openapi.yml"'+"\n")

    #build faucet
    file.write("faucet:"+"\n")
    file.write("  name: "+"household1"+"\n")
    file.write('  coins: ["5token", "100000stake"]'+"\n")
    file.close()

    #create file transaction.sh
    file = open("transaction.sh", "w")

    file.write("#!/bin/bash"+"\n")
    for household in range(int(households)):
         file.write("household"+str(household)+"=$(docker exec checkers checkersd keys show household"+str(household)+ " -a)"+"\n")
         file.write("export household"+str(household)+"\n")
         file.write('echo "household'+str(household)+': $household'+str(household)+'"'+"\n")