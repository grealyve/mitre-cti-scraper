package charts

import (
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/grealyve/mitre-cti-scraper/database"
	"github.com/grealyve/mitre-cti-scraper/models"
	"gorm.io/gorm"
)

var (
	databaseConnection *gorm.DB
	config             = &models.APTFeedConfigDB{
		Host:     "localhost",
		Port:     "5432",
		Password: "seferalgan",
		User:     "minibots",
		DBName:   "mitre_cti_db",
		SSLMode:  "disable",
	}

	platformCounters      = make(map[string]int)
	relationshipTypeCount = make(map[string]int)
	killChainPhasesMap    = make(map[string]int)
)

func init() {
	//Connect to the database
	db, err := database.ConnectAndMigrateDB(config)
	databaseConnection = db
	if err != nil {
		fmt.Errorf("Failed to connect to the database")
	}
}

func ScetchDiagrams() {
	// Pie chart scetch diagram
	pieChartPage := components.NewPage()
	pieChartPage.AddCharts(
		subTechnicPieChart(),
	)
	f, err := os.Create("charts/pie.html")
	if err != nil {
		fmt.Errorf("Failed to create file pie.html")
	}
	pieChartPage.Render(io.MultiWriter(f))

	// Bar chart scetch diagram
	barChartPage := components.NewPage()
	barChartPage.AddCharts(
		platformBarChart(),
		killChainPhasesBarChart(),
	)
	f, err = os.Create("charts/bar.html")
	if err != nil {
		fmt.Errorf("Failed to create file bar.html")
	}
	barChartPage.Render(io.MultiWriter(f))

	// Wordcloud chart scetch diagram
	wordCloudPage := components.NewPage()
	wordCloudPage.AddCharts(
		wordCloudPlatformChart(),
	)
	f, err = os.Create("charts/wordcloud.html")
	if err != nil {
		fmt.Errorf("Failed to create file wordcloud.html")
	}
	wordCloudPage.Render(io.MultiWriter(f))

	// Pie chart scetch diagram
	relationshipTypePieChartPage := components.NewPage()
	relationshipTypePieChartPage.AddCharts(
		relationshipTypePieChart(),
	)
	f, err = os.Create("charts/relationshipTypePie.html")
	if err != nil {
		fmt.Errorf("Failed to create file relationshipTypePie.html")
	}
	relationshipTypePieChartPage.Render(io.MultiWriter(f))

	killchainPhasesPieChartPage := components.NewPage()
	killchainPhasesPieChartPage.AddCharts(
		killChainPhasesPieChart(),
	)
	f, err = os.Create("charts/killChainPhasesPie.html")
	if err != nil {
		fmt.Errorf("Failed to create file killChainPhasesPie.html")
	}
	killchainPhasesPieChartPage.Render(io.MultiWriter(f))
}

func generateSubtechnicsPieItems() []opts.PieData {
	subtechnicCount := 0

	technicData, err := database.GetAPTFeedTechnicObjectDataList(databaseConnection)
	if err != nil {
		fmt.Errorf("Failed to get technic data")
	}

	for _, technic := range technicData {
		if technic.XMitreIsSubtechnique {
			subtechnicCount++
		}
	}

	items := make([]opts.PieData, 0)
	items = append(items, opts.PieData{Name: "Subtechnic", Value: subtechnicCount})
	items = append(items, opts.PieData{Name: "Main Technic", Value: (len(technicData) - subtechnicCount)})

	return items
}

func subTechnicPieChart() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Subtechnics Pie Chart"}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithAnimation(),
	)

	pie.AddSeries("pie", generateSubtechnicsPieItems()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
		)

	return pie
}

func generatePlatformBarItems() []opts.BarData {
	technicData, err := database.GetAPTFeedTechnicObjectDataList(databaseConnection)
	if err != nil {
		fmt.Errorf("Failed to get technic data")
	}

	for _, technic := range technicData {
		if technic.XMitrePlatforms != nil {
			for _, platform := range technic.XMitrePlatforms {
				platformCounters[platform]++
			}
		}
	}

	items := make([]opts.BarData, 0)
	for platform, count := range platformCounters {
		items = append(items, opts.BarData{Value: count, Name: platform})
	}

	return items
}

func getPlatforms() []string {
	platforms := make([]string, 0)
	for platform := range platformCounters {
		platforms = append(platforms, platform)
	}
	return platforms
}

func platformBarChart() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Platforms Bar Chart"}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Platforms"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Count"}),
		charts.WithAnimation(),
	)

	bar.SetXAxis(getPlatforms()).
		AddSeries("Platforms", generatePlatformBarItems()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
		)

	return bar
}

func generateWordCloudData(data map[string]int) (items []opts.WordCloudData) {
	items = make([]opts.WordCloudData, 0)
	for k, v := range data {
		items = append(items, opts.WordCloudData{Name: k, Value: v})
	}
	return
}

func wordCloudPlatformChart() *charts.WordCloud {
	wordCloud := charts.NewWordCloud()
	wordCloud.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Mitre Platforms WordCloud Chart",
		}))

	wordCloud.AddSeries("wordcloud", generateWordCloudData(platformCounters)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "cardioid",
				}),
		)

	return wordCloud
}

func generateRelationshipTypeItems() []opts.PieData {
	relationshipData, err := database.GetAPTFeedRelationshipObjectDataList(databaseConnection)
	if err != nil {
		fmt.Errorf("Failed to get relationship data")
	}

	for _, relationship := range relationshipData {
		relationshipTypeCount[relationship.RelationshipType]++
	}

	items := make([]opts.PieData, 0)
	for relationshipType, count := range relationshipTypeCount {
		items = append(items, opts.PieData{Name: relationshipType, Value: count})
	}

	return items
}

func relationshipTypePieChart() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Relationship Types Pie Chart"}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithAnimation(),
	)

	pie.AddSeries("pie", generateRelationshipTypeItems()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
		)

	return pie
}

func generateTechnicsKillChainPhasesItems() []opts.BarData {
	technicData, err := database.GetAPTFeedTechnicObjectDataList(databaseConnection)
	if err != nil {
		fmt.Errorf("Failed to get technic data")
	}

	for _, technic := range technicData {
		for _, killChainPhase := range technic.KillChainPhases {
			killChainPhasesMap[killChainPhase.PhaseName]++
		}
	}

	items := make([]opts.BarData, 0)
	for killChainPhase, count := range killChainPhasesMap {
		items = append(items, opts.BarData{Value: count, Name: killChainPhase})
	}

	return items
}

func killChainPhasesBarChart() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Kill Chain Phases Bar Chart"}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Kill Chain Phases"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Count"}),
		charts.WithAnimation(),
	)

	bar.SetXAxis(getKillChainPhases()).
		AddSeries("Kill Chain Phases", generateTechnicsKillChainPhasesItems()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
		)

	return bar
}

func getKillChainPhases() []string {
	killChainPhases := make([]string, 0)
	for killChainPhase := range killChainPhasesMap {
		killChainPhases = append(killChainPhases, killChainPhase)
	}
	return killChainPhases
}

func killChainPhasesPieChart() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Kill Chain Phases Pie Chart"}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithAnimation(),
	)

	pie.AddSeries("pie", generateTechnicsKillChainPhasesItemsPieChart()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
		)

	return pie
}

func generateTechnicsKillChainPhasesItemsPieChart() []opts.PieData {
	technicData, err := database.GetAPTFeedTechnicObjectDataList(databaseConnection)
	if err != nil {
		fmt.Errorf("Failed to get technic data")
	}

	for _, technic := range technicData {
		for _, killChainPhase := range technic.KillChainPhases {
			killChainPhasesMap[killChainPhase.PhaseName]++
		}
	}

	items := make([]opts.PieData, 0)
	for killChainPhase, count := range killChainPhasesMap {
		items = append(items, opts.PieData{Name: killChainPhase, Value: count})
	}

	return items

}
