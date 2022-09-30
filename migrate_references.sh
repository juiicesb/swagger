flist=`grep -r "github.com/juiicesb/swagger" * | awk -F ':' '{print $1}' | sort | uniq`
for f in $flist; do echo $f; sed -i -e "s|github.com/juiicesb/swagger|github.com/juiicesb/swagger|g" $f; done


flist=`grep -r "saas-kit" * | awk -F ':' '{print $1}' | sort | uniq`
for f in $flist; do echo $f; sed -i -e "s|swagger|swagger|g" $f; done


flist=`grep -r "saas-kit" * | awk -F ':' '{print $1}' | sort | uniq`
for f in $flist; do echo $f; sed -i -e "s|swagger|swagger|g" $f; done
