define(['alertify'],function(alertify){
	$(document).on('click','.delete',function(e){
		var that = this;
		alertify.confirm('Are you sure you want to delete this content?',function(e){
			var id = $(that).data('id');
			$.ajax({
				url: '/admin/equipment/'+id,
				type: 'DELETE',
				success:function(data, status, xhr){
					$(that).closest('tr').remove();
					alertify.success('Content removed');
				},
				error: function(xhr, status, err){
					alertify.error(xhr.responseText);
				}
			});
		});
		return false;
	}).on('click','.add-component', function(){
		var comp = $('.new-component').val();
		if (comp.length === 0){
			alertify.error('Component cannot be empty');
			return false;
		}
		var idx = $('.list-group-item').length;
		var html =
						'<li class="list-group-item">'+
						'<input type="hidden" name="component['+idx+']" value="'+comp+'">'+
						comp+
						'<span class="glyphicon glyphicon-remove-circle pull-right remove-component"></span>' +
						'</li>';
		$('.component-list').append(html);
		$('.new-component').val('');
		return false;
	}).on('click','.remove-component',function(){
		$(this).closest('li').remove();
	});
});