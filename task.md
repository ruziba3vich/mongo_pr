MongoDB Uyga Vazifa Topshiriqlari
Ushbu uyga vazifa topshiriqlari MongoDB dan foydalanishni o'rganishga yordam beradi. Har bir topshiriqni bajarishdan oldin, kerakli dasturiy ta'minotni o'rnatishingiz va MongoDB serverini ishga tushirishingiz kerak.
1. Dastlabki shart
Topshiriq:
Shu link boyicha Query Document mavzusidan Limiting Records ga cha mavzularini k'orib chiqing
Yangi home tasks nomli database yarating
tasks nomli collection qo'shing
2. Hujjatlar Qo'shish
Topshiriq:
Quyidagi hujjatlarni bitta buyruq bilan qo'shing:
{ 
    title: "Oziq-ovqat sotib olish", 
    description: "Sut, Non, Pishloq", 
    status: "pending", 
    assignedTo: { name: "John Doe", email: "john@example.com" }, 
    dueDate: new Date("2024-07-10"), 
    subTasks: [
        { title: "Sut sotib olish", status: "pending" },
        { title: "Non sotib olish", status: "pending" }
    ]
},
{ 
    title: "Uy vazifasini tugatish", 
    description: "Matematika va Fan", 
    status: "in-progress", 
    assignedTo: { name: "Jane Smith", email: "jane@example.com" }, 
    dueDate: new Date("2024-07-12"), 
    subTasks: [
        { title: "Matematika vazifasi", status: "in-progress" },
        { title: "Fan loyihasi", status: "boshlanmagan" }
    ]
},
{ 
    title: "Uy tozalash", 
    description: "Yashash xonasi va oshxona", 
    status: "pending", 
    assignedTo: { name: "Alice Brown", email: "alice@example.com" }, 
    dueDate: new Date("2024-07-15"), 
    subTasks: [
        { title: "Yashash xonasini tozalash", status: "pending" },
        { title: "Oshxonani tozalash", status: "pending" }
    ]
}

3. Vazifalarni topish
Topshiriq:
Muayyan foydalanuvchiga yuklangan barcha vazifalarni toping
Muayyan sanadan oldin tugash muddati bo'lgan vazifalarni toping
Tugatilmagan qo'shimcha vazifalari bo'lgan vazifalarni toping
4. Vazifalarni yangilash
Topshiriq:
Qo'shimcha (sub-tasks) vazifaning holatini yangilang
Ihtiyoriy vazifani boshqa foydalanuvchiga o'tkazing
Mavjud vazifaga yangi qo'shimcha vazifa qo'shing
5. Vazifalarni o'chirish
Topshiriq:
Muayyan qo'shimcha vazifani o'chiring
Muayyan foydalanuvchiga topshirilgan barcha vazifalarni o'chiring
O'tmishda tugatish muddati bo'lgan vazifalarni o'chiring
Qo'shimcha talab
Har bir vazifa uchun bajarilgan komandalarni va natijalarni ko'rsatadigan skrinshot qo'shing
