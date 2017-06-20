new Vue({
  el: '#app',

  data: {
    input: '',
    encodeds: [],
    loading: false,
  },

  computed: {
    trimmedInput() {
      return this.input.trim()
    },
    noResult() {
      return !this.loading && this.trimmedInput !== '' && this.encodeds.length === 0
    },
    alphabet() {
      return this.encodeds.length ? this.trimmedInput : 'alphabet'
    },
    quran() {
      return this.encodeds.length ? this.encodeds[0].text.Enhanced : "Al-Qu'ran"
    },
  },

  watch: {
    input: function(newInput) {
      this.updateResult()
    },
  },

  methods: {
    updateResult: _.debounce(function() {
      this.loading = true
      axios.get('/api/encode/' + this.trimmedInput)
      .then((response) => this.encodeds = response.data.map((text) => ({text})))
      .catch(() => this.encodeds = [])
      .then(() => this.loading = false)
    }, 500),

    locate(encoded) {
      this.$set(encoded, 'expanded', !encoded.expanded)
      if (encoded.locations) return
      this.$set(encoded, 'loading', true)
      axios.get('/api/locate/' + encoded.text.Clean)
      .then((response) => {
        let data = response.data
        data.forEach((location) => {
          let aya = location.Aya
          let begin = location.Location.Index
          let end = location.Location.Index + encoded.text.Clean.length
          location.AyaBeforeHL = aya.substring(0, begin)
          location.AyaHL = aya.substring(begin, end)
          location.AyaAfterHL = aya.substring(end)
        })
        this.$set(encoded, 'locations', data)
      })
      .catch(() => this.$set(encoded, 'locations', undefined))
      .then(() => this.$set(encoded, 'loading', false))
    },
  },
})