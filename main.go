package main

import (
	"bufio"
	"fmt"
	"go-training-gof-design-pattern/abstract_factory"
	"go-training-gof-design-pattern/adapter"
	"go-training-gof-design-pattern/bridge"
	"go-training-gof-design-pattern/builder"
	"go-training-gof-design-pattern/chain_of_responsibility"
	"go-training-gof-design-pattern/command"
	"go-training-gof-design-pattern/composite"
	"go-training-gof-design-pattern/decorator"
	"go-training-gof-design-pattern/factory_method"
	"go-training-gof-design-pattern/flyweight"
	"go-training-gof-design-pattern/interpreter"
	"go-training-gof-design-pattern/iterator"
	"go-training-gof-design-pattern/mediator"
	"go-training-gof-design-pattern/memento"
	"go-training-gof-design-pattern/observer"
	"go-training-gof-design-pattern/prototype"
	"go-training-gof-design-pattern/proxy"
	"go-training-gof-design-pattern/singleton"
	"go-training-gof-design-pattern/state"
	"go-training-gof-design-pattern/strategy"
	"go-training-gof-design-pattern/template_method"
	"go-training-gof-design-pattern/visitor"
	"os"
	"time"
)

// main
// いちいち入れ替えて実行してみているので各パターンのテストケースとして実行ファイルを配置したほうがいい。
func main() {
	runInterpreter()
}

func runInterpreter() {
	file, err := os.Open("interpreter/program.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("text = \"%s\"\n", text)
		node := interpreter.NewProgramNode()
		err := node.Parse(interpreter.NewContext(text))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("node = %s", node.ToString())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func runCommand() {
	// Client ConcreteCommandを生成し、Receiverにを割り当てる
	history := command.NewMacroCommand()
	canvas := command.NewDrawCanvas(history) // Receiver
	cmd1 := command.NewDrawCommand(canvas, command.NewPosition(1, 1))
	cmd2 := command.NewDrawCommand(canvas, command.NewPosition(1, 2))
	history.Append(cmd1)
	history.Append(cmd2)

	// Invoker コマンドの実行者
	func(cmd command.Command) {
		history.Execute()
		history.Undo()
		history.Execute()
	}(history)
}

func runProxy() {
	p := proxy.NewPrinterProxy("Alice")
	func(p proxy.Printable) {
		fmt.Printf("名前は現在%sです。\n", p.GetPrinterName())
		p.SetPrinterName("Bob")
		fmt.Printf("名前は現在%sです。\n", p.GetPrinterName())
		p.Print("Hello, World.")
	}(p)
}

func runFlyweight() {
	bs := flyweight.NewBigString("124")
	bs.Print()
}

func runState() {
	frame := state.NewSafeFrame("State Sample")
	for i := 0; i < 24; i++ {
		frame.SetClock(i)
		// もとのサンプルは画面コンポーネントのリスナーが呼び出してるけどそれはないので擬似的にmainから操作として呼び出す
		frame.State.DoUse(frame)
	}
	// 画面がないので標準出力にて状態ごとの処理が別れていることを確認する。
	fmt.Printf("%+v\n", frame.TextScreen)
}

func runMoment() {
	gamer := memento.NewGamer(100)
	m := gamer.CreateMemento()
	for i := 0; i < 100; i++ {
		fmt.Printf("==== %d\n", i)
		fmt.Printf("現状:%s", gamer)
		gamer.Bet()
		fmt.Printf("所持金は%dになりました。\n", gamer.Money())
		if gamer.Money() > m.Money() {
			fmt.Println("    だいぶ増えたので、現在の状態を保存しておこう。")
			m = gamer.CreateMemento()
		} else if gamer.Money() < m.Money() {
			fmt.Println("    だいぶ減ったので、以前の状態に復帰しよう。")
			gamer.RestoreMemento(m)
		}
		time.Sleep(time.Second * 1)
	}
}

func runObserver() {
	generator := observer.NewIncrementalNumberGenerator(10, 50, 5)
	digitObserver := observer.NewDigitObserver()
	graphObserver := observer.NewGraphObserver()
	generator.AddObserver(digitObserver)
	generator.AddObserver(graphObserver)
	generator.Execute()
}

func runMediator() {
	frame := mediator.NewLoginFrame("Title")
	frame.Print()
	frame.CheckGuest.SetState(false)
	frame.CheckGuest.ItemStateChanged()
	frame.Print()
	frame.TextUser.SetText("natacon")
	frame.TextUser.TextValueChanged()
	frame.Print()
	frame.TextPass.SetText("natacon")
	frame.TextPass.TextValueChanged()
	frame.Print()
}

func runChainOfResponsibility() {
	alice := chain_of_responsibility.NewNoSupport("Alice")
	bob := chain_of_responsibility.NewLimitSupport("Bob", 100)
	charlie := chain_of_responsibility.NewSpecialSupport("Charlie", 429)
	diana := chain_of_responsibility.NewLimitSupport("Diana", 200)
	elmo := chain_of_responsibility.NewOddSupport("Elmo")
	fred := chain_of_responsibility.NewLimitSupport("Fred", 300)
	alice.
		SetNext(bob).
		SetNext(charlie).
		SetNext(diana).
		SetNext(elmo).
		SetNext(fred)
	for i := 0; i < 500; i += 33 {
		alice.Support(chain_of_responsibility.NewTrouble(i))
	}
}

func runVisitor() {
	fmt.Println("Making root entries...")
	rootdir := visitor.NewDirectory("root")
	bindir := visitor.NewDirectory("bin")
	tmpdir := visitor.NewDirectory("tmp")
	usrdir := visitor.NewDirectory("usr")
	rootdir.Add(bindir)
	rootdir.Add(tmpdir)
	rootdir.Add(usrdir)
	bindir.Add(visitor.NewFile("vi", 10000))
	bindir.Add(visitor.NewFile("latex", 20000))
	//rootdir.Accept(visitor.NewListVisitor())
	tmpdir.Add(visitor.NewFile("visitor.html", 10))
	rootdir.Accept(visitor.NewListVisitor())
	fmt.Println("---")
	ffv := visitor.NewFileFindVisitor(".html")
	rootdir.Accept(ffv)
	for _, entry := range ffv.FoundFiles() {
		fmt.Println(entry)
	}
}
func runDecorator() {
	b1 := decorator.NewStringDisplay("Hello World")
	b2 := decorator.NewSideBorder(b1, "#")
	b3 := decorator.NewFullBorder(b2)
	b4 := decorator.NewUpDownBorder(b1, "!")
	b1.Show()
	b2.Show()
	b3.Show()
	b4.Show()

	fmt.Println("---")
	b5 := decorator.NewMultiStringDisplay()
	b5.Add("おはよう")
	b5.Add("こんにちは")
	b6 := decorator.NewSideBorder(b5, "#")
	b6.Show()
}

func runComposite() {
	fmt.Println("Making root entries...")
	rootDir := composite.NewDirectoryEntry("root")
	binDir := composite.NewDirectoryEntry("bin")
	tmpDir := composite.NewDirectoryEntry("tmp")
	usrDir := composite.NewDirectoryEntry("usr")
	rootDir.Add(binDir)
	rootDir.Add(tmpDir)
	rootDir.Add(usrDir)
	binDir.Add(composite.NewFileEntry("vi", 10000))
	binDir.Add(composite.NewFileEntry("latex", 20000))
	rootDir.PrintList("")
}

func runStrategy() {
	player1 := strategy.NewPlayer("Taro", strategy.NewWinningStrategy())
	player2 := strategy.NewPlayer("Hana", strategy.NewRandomStrategy())
	for i := 0; i < 10000; i++ {
		nextHand1 := player1.NextHand()
		nextHand2 := player2.NextHand()
		if nextHand1.IsStrongerThan(nextHand2) {
			fmt.Printf("Winner:%s\n", player1)
			player1.Win()
			player2.Lose()
		} else if nextHand1.ISWeakerThan(nextHand2) {
			fmt.Printf("Winner:%s\n", player2)
			player1.Lose()
			player2.Win()
		} else {
			fmt.Println("Even...")
			player1.Even()
			player2.Even()
		}
	}
	fmt.Printf("Total Result:\n%s\n%s\n", player1, player2)
}

func runBridge() {
	display1 := bridge.NewDisplay(bridge.NewStringDisplay("Hello Japan!"))
	display2 := bridge.NewDisplay(bridge.NewStringDisplay("Hello World!"))
	display3 := bridge.NewCountDisplay(bridge.NewStringDisplay("Hello, Universe!"))
	display4 := bridge.NewRandomDisplay(bridge.NewStringDisplay("Hello Random!"))
	display5 := bridge.NewIncreaseDisplay(bridge.NewCharDisplayImpl('a', 'b', 'c'), 5)
	display1.DisplayFunc()
	display2.DisplayFunc()
	display3.DisplayFunc()
	display3.MultiDisplay(5)
	display4.RandomDisplay(5)
	display4.MultiDisplay(5)
	display5.IncreaseDisplay(5)

}

func runAbstractFactory() {
	listFactory := abstract_factory.NewListFactory()
	yahoo := listFactory.CreateLink("Yahoo", "http://www.yahoo.co.jp/")
	google := listFactory.CreateLink("Google", "http://www.google.com/")

	traySearch := listFactory.CreateTray("サーチエンジン")
	traySearch.Add(yahoo)
	traySearch.Add(google)

	page := listFactory.CreatePage("LinkPage", "natacon")
	page.Add(traySearch)
	page.Output(page)
}

func runBuilder() {
	textBuilder := builder.NewTextBuilder()
	director := builder.NewDirector(textBuilder)
	director.Construct()
	fmt.Println(textBuilder.Result())

	htmlBuilder := builder.NewHTMLBuilder()
	director2 := builder.NewDirector(htmlBuilder)
	director2.Construct()
	fmt.Println(htmlBuilder.Result())
}

func runPrototype() {
	manager := prototype.NewManager()
	upen := prototype.NewUnderlinePen('~')
	mbox := prototype.NewMessageBox('*')
	sbox := prototype.NewMessageBox('/')
	manager.Register("strong message", upen)
	manager.Register("warning box", mbox)
	manager.Register("slash box", sbox)

	p1 := manager.Create("strong message")
	p1.Use("Hello, world.")
	p2 := manager.Create("warning box")
	p2.Use("Hello, world.")
	p3 := manager.Create("slash box")
	p3.Use("Hello, world.")
}

func runSingleton() {
	instance1 := singleton.GetInstance()
	instance1.SetI(1)
	instance2 := singleton.GetInstance()
	instance2.SetI(2)
	instance3 := singleton.GetInstance()
	instance3.SetI(3)
	// iが同じ値であることを確認する
	fmt.Println(instance1, instance2, instance3)
}

func runFactoryMethod() {
	factory := factory_method.NewIDCardFactory()
	card1 := factory.Create("natacon1")
	card1.Use()
	fmt.Println(factory.Owners())
}

func runTemplateMethod() {
	display1 := template_method.NewCharDisplay('A')
	display1.Display()
	display2 := template_method.NewStringDisplay("Hello World!")
	display2.Display()
}

func runAdapter() {
	p := adapter.NewPrintBanner("Adapter!")
	p.PrintWeak()
	p.PrintStrong()
}

func runIterator() {
	bookShelf := iterator.NewBookShelf()
	bookShelf.Append(iterator.NewBook("A"))
	bookShelf.Append(iterator.NewBook("B"))
	bookShelf.Append(iterator.NewBook("C"))
	bookShelf.Append(iterator.NewBook("D"))

	// TODO: Javaから愚直にやるとこんなんなるから考え方が違うはず。
	it := bookShelf.Iterator()
	for it.HasNext() {
		book, ok := it.Next().(iterator.Book)
		if !ok {
			fmt.Printf("Failed to type assertion: %v\n", book)
			continue
		}
		fmt.Println(book)
	}
}
