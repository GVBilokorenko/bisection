helpFunction()
{
   echo "Usage: bisection -l \"float\" -u \"float\" -a \"float\" -c \"string\""
   echo -e "\t-l Lower bound"
   echo -e "\t-u Upper bound"
   echo -e "\t-a Accuracy"
   echo -e "\t-c Coefficients of polynom; example: \"1.0 -1.0 0.0 2\""
   exit 1 # Exit script after printing help
}

if [ "$1" == "-h" ]
then
    helpFunction
fi

while getopts "l:u:a:c:" opt
do
    case "$opt" in
        l ) parameterL="$OPTARG" ;;
        u ) parameterU="$OPTARG" ;;
        a ) parameterA="$OPTARG" ;;
        c ) parameterC="$OPTARG" ;;
        ? ) helpFunction ;; # Print helpFunction in case parameter is non-existent
    esac
done

if [ -z "$parameterL" ] || [ -z "$parameterU" ] || [ -z "$parameterA" ] || [ -z "$parameterC" ]
then
   echo "Some or all of the parameters are empty";
   helpFunction
fi
