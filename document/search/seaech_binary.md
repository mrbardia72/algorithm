<p dir="rtl">
در این روش مرتب بودن مهم هست(اگر مرتب نباشد غیرقابل استفاده هست) و معمولا به صورت پیش فرض صعودی هست
 <br>
 <br>
روش جست و جوی این الگوریتم
 <br>
 <br>
در جستجوی دودویی (Binary Search)‌، جستجو در یک آرایه مرتب شده، با تقسیم تکرار شونده بازه جستجو به نصف، انجام می‌شود. کار با بازه‌ای که کل آرایه را پوشش می‌دهد، آغاز می‌شود. اگر مقدار کلید جستجو برابر با عنصر میانی باشد، اندیس آن بازگردانده می‌شود.
 <br>
 <br>
در غیر این صورت، اگر مقدار کلید جستجو کمتر از عنصری باشد که در میانه بازه قرار دارد، بازه شکسته شده و جستجو در نیمه کمتر ادامه پیدا می‌کند. در صورتی که مقدار کلید جستجو بزرگ‌تر از اندیس میانی آرایه باشد، جستجو در نیمه بیشتر (حاوی مقادیر بزرگ‌تر) آرایه ادامه پیدا می‌کند. کار شکستن آرایه به دو نیم و انتخاب نیمه‌ای که جستجو باید در آن انجام شود، مکررا و تا هنگامی که عنصر مورد نظر در آرایه یافته شود و یا مشخص شود که عنصر مورد نظر در آرایه وجود ندارد، ادامه خواهد داشت.
 <br>
 <br>
جستجوی موفق: حداقل تعداد مقایسه یک می باشد
 <br>
جستجوی موفق: حداکثر تعداد مقایسه 1+[log2 n] می باشد
 <br>
جستجوی موفق: اگر عنصر موردنظر در آرایه ما موجود باشد همواره با حداکثر (log2 n)o پیدا می شود
 <br>
جستجوی موفق: برای دو مقدار متفاوت تعداد مقایسه های لازم لزوما برابر نیست
 <br>
 <br>
جستجوی ناموفق: حداقل تعداد مقایسه [log2 n] می باشد
<br>
جستجوی ناموفق: حداکثر تعداد مقایسه 1+[log2 n] می باشد
<br>
جستجوی ناموفق: اگر عنصر موردنظر در آرایه ما موجود نباشد همواره با حداکثر (log2 n)o پیدا می شود
<br>
جستجوی ناموفق: برای دو مقدار متفاوت تعداد مقایسه های لازم لزوما برابر نیست
<br>
<br>
<br>
 متوسط تعداد مقایسه های جستجوی موفق s(n) می باشد
<br>
s(n)=(Total successful comparisons/n)
<br>
<br>
 متوسط تعداد مقایسه های جستجوی ناموفق u(n) می باشد
<br>
u(n)=(Total failed comparisons/n)
<br>
<br>
بهترین حالت این روش حداقا مقایسه یک در زمان یک می باشد
<br>
حالت متوسط: کمتر از 1+[log2 n] در زمان (log2 n)o می باشد
<br>
حالت بدترین: حداکثر مقایسه  1+[log2 n] در زمان (log2 n)o می باشد
<br>
<br>
تعداد مقایسه برای یافتن کوچکترین و بزرگترین عدد 3n/2-2 می باشد
<br>
</p>