flist=`grep -r "github.com/juiicesb/swaggergerger" * | awk -F ':' '{print $1}' | sort | uniq`
for f in $flist; do echo $f; sed -i -e "s|github.com/juiicesb/swaggergerger|github.com/juiicesb/swaggergerger|g" $f; done


flist=`grep -r "saas-kit" * | awk -F ':' '{print $1}' | sort | uniq`
for f in $flist; do echo $f; sed -i -e "s|swaggerger|swaggergerger|g" $f; done


flist=`grep -r "saas-kit" * | awk -F ':' '{print $1}' | sort | uniq`
for f in $flist; do echo $f; sed -i -e "s|swaggerger|swaggergerger|g" $f; done
